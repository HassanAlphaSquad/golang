<title>PHP</title>
<?php
// echo "Hello<br/>";
// $types = [
//     1 => "text",
//     2 => "video",
//     3 => "audio",
//     4 => "image"
// ];

// foreach ($types as $key => $value) {
//     echo "$key => $value<br>";
// }



// -------- Variables & Data Types --------
$string = "Hello, PHP!";
$integer = 42;
$float = 3.14;
$boolean = true;
$array = ["Apple", "Banana", "Cherry"];

echo "<h3>Variables & Data Types</h3>";
echo "<p>String: $string</p>";
echo "<p>Integer: $integer</p>";
echo "<p>Float: $float</p>";
echo "<p>Boolean: " . ($boolean ? "true" : "false") . "</p>";
echo "<p>Array: " . implode(", ", $array) . "</p>";

// -------- Constants --------
define("PI", 3.14159);
echo "<h3>Constants</h3>";
echo "<p>PI Value: " . PI . "</p>";

// -------- Control Structures (if-else, loops) --------
echo "<h3>Control Structures</h3>";
if ($integer > 20) {
    echo "<p>$integer is greater than 20</p>";
} else {
    echo "<p>$integer is not greater than 20</p>";
}

echo "<h4>Looping through an array:</h4>";
foreach ($array as $fruit) {
    echo "<p>$fruit</p>";
}

// -------- Functions --------
function greet($name = "Guest")
{
    return "Hello, $name!";
}
echo "<h3>Functions</h3>";
echo "<p>" . greet("John") . "</p>";

// -------- Classes & Objects --------
class Car
{
    public $brand;
    public $color;

    public function __construct($brand, $color)
    {
        $this->brand = $brand;
        $this->color = $color;
    }

    public function getDetails()
    {
        return "This is a $this->color $this->brand.";
    }
}

$myCar = new Car("Toyota", "Red");
echo "<h3>OOP: Classes & Objects</h3>";
echo "<p>" . $myCar->getDetails() . "</p>";

// -------- Error Handling --------
echo "<h3>Error Handling</h3>";
function divide($a, $b)
{
    if ($b == 0) {
        throw new Exception("Division by zero!");
    }
    return $a / $b;
}

try {
    echo "<p>10 / 2 = " . divide(10, 2) . "</p>";
    echo "<p>10 / 0 = " . divide(10, 0) . "</p>"; // This will cause an exception
} catch (Exception $e) {
    echo "<p>Error: " . $e->getMessage() . "</p>";
}

// -------- File Handling (Writing & Reading) --------
echo "<h3>File Handling</h3>";
$filename = "sample.txt";
file_put_contents($filename, "Hello, this is a test file!");

if (file_exists($filename)) {
    $content = file_get_contents($filename);
    echo "<p>File Content: $content</p>";
} else {
    echo "<p>File not found!</p>";
}
?>