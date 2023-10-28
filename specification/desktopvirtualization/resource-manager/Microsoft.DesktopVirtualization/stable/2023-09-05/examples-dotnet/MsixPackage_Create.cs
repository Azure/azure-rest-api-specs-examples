using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DesktopVirtualization;
using Azure.ResourceManager.DesktopVirtualization.Models;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/MsixPackage_Create.json
// this example is just showing the usage of "MSIXPackages_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HostPoolResource created on azure
// for more information of creating HostPoolResource, please refer to the document of HostPoolResource
string subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
string resourceGroupName = "resourceGroup1";
string hostPoolName = "hostpool1";
ResourceIdentifier hostPoolResourceId = HostPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hostPoolName);
HostPoolResource hostPool = client.GetHostPoolResource(hostPoolResourceId);

// get the collection of this MsixPackageResource
MsixPackageCollection collection = hostPool.GetMsixPackages();

// invoke the operation
string msixPackageFullName = "msixpackagefullname";
MsixPackageData data = new MsixPackageData()
{
    ImagePath = "imagepath",
    PackageName = "MsixPackage_name",
    PackageFamilyName = "MsixPackage_FamilyName",
    DisplayName = "displayname",
    PackageRelativePath = "packagerelativepath",
    IsRegularRegistration = false,
    IsActive = false,
    PackageDependencies =
    {
    new MsixPackageDependencies()
    {
    DependencyName = "MsixTest_Dependency_Name",
    Publisher = "PublishedName",
    MinVersion = "version",
    }
    },
    Version = "version",
    LastUpdatedOn = DateTimeOffset.Parse("2008-09-22T14:01:54.9571247Z"),
    PackageApplications =
    {
    new MsixPackageApplications()
    {
    AppId = "ApplicationId",
    Description = "application-desc",
    AppUserModelId = "AppUserModelId",
    FriendlyName = "friendlyname",
    IconImageName = "Apptile",
    RawIcon = BinaryData.FromString("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
    RawPng = BinaryData.FromString("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
    }
    },
};
ArmOperation<MsixPackageResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, msixPackageFullName, data);
MsixPackageResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MsixPackageData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
