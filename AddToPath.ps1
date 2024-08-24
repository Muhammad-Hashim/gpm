# Path to the directory containing gpm.exe
$gpmDirectory = "C:\Tools"

# Get the current PATH environment variable
$envPath = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::Machine)

# Check if the directory is already in the PATH
if ($envPath -notlike "*$gpmDirectory*") {
    # Add the directory to PATH
    $newPath = "$envPath;$gpmDirectory"
    [System.Environment]::SetEnvironmentVariable("Path", $newPath, [System.EnvironmentVariableTarget]::Machine)
    Write-Output "Directory added to PATH: $gpmDirectory"
} else {
    Write-Output "Directory is already in PATH: $gpmDirectory"
}
