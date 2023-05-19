using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DesktopVirtualization;
using Azure.ResourceManager.DesktopVirtualization.Models;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2022-09-09/examples/MsixPackage_Delete.json
// this example is just showing the usage of "MSIXPackages_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MsixPackageResource created on azure
// for more information of creating MsixPackageResource, please refer to the document of MsixPackageResource
string subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
string resourceGroupName = "resourceGroup1";
string hostPoolName = "hostpool1";
string msixPackageFullName = "packagefullname";
ResourceIdentifier msixPackageResourceId = MsixPackageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hostPoolName, msixPackageFullName);
MsixPackageResource msixPackage = client.GetMsixPackageResource(msixPackageResourceId);

// invoke the operation
await msixPackage.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
