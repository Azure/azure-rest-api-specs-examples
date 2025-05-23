using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DesktopVirtualization.Models;
using Azure.ResourceManager.DesktopVirtualization;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/MsixPackage_Update.json
// this example is just showing the usage of "MSIXPackages_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MsixPackageResource created on azure
// for more information of creating MsixPackageResource, please refer to the document of MsixPackageResource
string subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
string resourceGroupName = "resourceGroup1";
string hostPoolName = "hostpool1";
string msixPackageFullName = "msixpackagefullname";
ResourceIdentifier msixPackageResourceId = MsixPackageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hostPoolName, msixPackageFullName);
MsixPackageResource msixPackage = client.GetMsixPackageResource(msixPackageResourceId);

// invoke the operation
MsixPackagePatch patch = new MsixPackagePatch
{
    IsActive = true,
    IsRegularRegistration = false,
    DisplayName = "displayname",
};
MsixPackageResource result = await msixPackage.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MsixPackageData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
