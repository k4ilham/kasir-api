# Test All API Endpoints
# Kasir API - Comprehensive Testing Script

Write-Host "==========================================" -ForegroundColor Cyan
Write-Host "  KASIR API - ENDPOINT TESTING" -ForegroundColor Cyan
Write-Host "==========================================" -ForegroundColor Cyan
Write-Host ""

$baseUrl = "http://localhost:8080"
$testResults = @()

function Test-Endpoint {
    param(
        [string]$Name,
        [string]$Method,
        [string]$Url,
        [string]$Body = $null
    )
    
    Write-Host "Testing: $Name" -ForegroundColor Yellow
    Write-Host "  Method: $Method" -ForegroundColor Gray
    Write-Host "  URL: $Url" -ForegroundColor Gray
    
    try {
        if ($Body) {
            $response = Invoke-WebRequest -Uri $Url -Method $Method -Body $Body -ContentType "application/json" -ErrorAction Stop
        } else {
            $response = Invoke-WebRequest -Uri $Url -Method $Method -ErrorAction Stop
        }
        
        $statusCode = $response.StatusCode
        $content = $response.Content
        
        Write-Host "  Status: $statusCode" -ForegroundColor Green
        Write-Host "  Response: $content" -ForegroundColor Green
        Write-Host ""
        
        return @{
            Name = $Name
            Status = "PASS"
            StatusCode = $statusCode
            Response = $content
        }
    }
    catch {
        Write-Host "  Status: FAILED" -ForegroundColor Red
        Write-Host "  Error: $($_.Exception.Message)" -ForegroundColor Red
        Write-Host ""
        
        return @{
            Name = $Name
            Status = "FAIL"
            Error = $_.Exception.Message
        }
    }
}

# 1. SYSTEM ENDPOINTS
Write-Host "=== SYSTEM ENDPOINTS ===" -ForegroundColor Cyan
$testResults += Test-Endpoint -Name "Home" -Method "GET" -Url "$baseUrl/"
$testResults += Test-Endpoint -Name "Health Check" -Method "GET" -Url "$baseUrl/health"

# 2. CATEGORIES ENDPOINTS
Write-Host "=== CATEGORIES ENDPOINTS ===" -ForegroundColor Cyan
$testResults += Test-Endpoint -Name "Get All Categories" -Method "GET" -Url "$baseUrl/api/categories"
$testResults += Test-Endpoint -Name "Get Category by ID" -Method "GET" -Url "$baseUrl/api/categories/1"

# Create new category
$newCategory = @{
    name = "Snack"
    description = "Kategori snack dan cemilan"
} | ConvertTo-Json

$testResults += Test-Endpoint -Name "Create Category" -Method "POST" -Url "$baseUrl/api/categories" -Body $newCategory

# Update category (ID 3)
$updateCategory = @{
    name = "Elektronik Updated"
    description = "Kategori produk elektronik dan gadget"
} | ConvertTo-Json

$testResults += Test-Endpoint -Name "Update Category" -Method "PUT" -Url "$baseUrl/api/categories/3" -Body $updateCategory

# 3. PRODUCTS ENDPOINTS
Write-Host "=== PRODUCTS ENDPOINTS ===" -ForegroundColor Cyan
$testResults += Test-Endpoint -Name "Get All Products" -Method "GET" -Url "$baseUrl/api/produk"
$testResults += Test-Endpoint -Name "Get Product by ID" -Method "GET" -Url "$baseUrl/api/produk/1"

# Create new product
$newProduct = @{
    name = "Laptop Asus"
    price = 7000000
    stock = 5
    category_id = 3
} | ConvertTo-Json

$testResults += Test-Endpoint -Name "Create Product" -Method "POST" -Url "$baseUrl/api/produk" -Body $newProduct

# Update product (ID 1)
$updateProduct = @{
    name = "Indomie Goreng Special"
    price = 4000
    stock = 15
    category_id = 1
} | ConvertTo-Json

$testResults += Test-Endpoint -Name "Update Product" -Method "PUT" -Url "$baseUrl/api/produk/1" -Body $updateProduct

# 4. SUMMARY
Write-Host "==========================================" -ForegroundColor Cyan
Write-Host "  TEST SUMMARY" -ForegroundColor Cyan
Write-Host "==========================================" -ForegroundColor Cyan
Write-Host ""

$passCount = ($testResults | Where-Object { $_.Status -eq "PASS" }).Count
$failCount = ($testResults | Where-Object { $_.Status -eq "FAIL" }).Count
$totalCount = $testResults.Count

Write-Host "Total Tests: $totalCount" -ForegroundColor White
Write-Host "Passed: $passCount" -ForegroundColor Green
Write-Host "Failed: $failCount" -ForegroundColor Red
Write-Host ""

if ($failCount -eq 0) {
    Write-Host "ALL TESTS PASSED!" -ForegroundColor Green
} else {
    Write-Host "SOME TESTS FAILED!" -ForegroundColor Red
    Write-Host ""
    Write-Host "Failed Tests:" -ForegroundColor Red
    $testResults | Where-Object { $_.Status -eq "FAIL" } | ForEach-Object {
        Write-Host "  - $($_.Name): $($_.Error)" -ForegroundColor Red
    }
}

Write-Host ""
Write-Host "==========================================" -ForegroundColor Cyan
