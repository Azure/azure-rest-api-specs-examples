using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DesktopVirtualization.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DesktopVirtualization;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/AppAttachPackage_Create.json
// this example is just showing the usage of "AppAttachPackage_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
string resourceGroupName = "resourceGroup1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this AppAttachPackageResource
AppAttachPackageCollection collection = resourceGroupResource.GetAppAttachPackages();

// invoke the operation
string appAttachPackageName = "msixpackagefullname";
AppAttachPackageData data = new AppAttachPackageData(new AzureLocation("southcentralus"), new AppAttachPackageProperties
{
    Image = new AppAttachPackageInfoProperties
    {
        PackageAlias = "msixpackagealias",
        ImagePath = "imagepath",
        PackageName = "MsixPackageName",
        PackageFamilyName = "MsixPackage_FamilyName",
        PackageFullName = "MsixPackage_FullName",
        DisplayName = "displayname",
        PackageRelativePath = "packagerelativepath",
        IsRegularRegistration = false,
        IsActive = false,
        PackageDependencies = {new MsixPackageDependencies
        {
        DependencyName = "MsixPackage_Dependency_Name",
        Publisher = "MsixPackage_Dependency_Publisher",
        MinVersion = "packageDep_version",
        }},
        Version = "packageversion",
        LastUpdatedOn = DateTimeOffset.Parse("2008-09-22T14:01:54.9571247Z"),
        PackageApplications = {new MsixPackageApplications
        {
        AppId = "AppId",
        Description = "PackageApplicationDescription",
        AppUserModelId = "AppUserModelId",
        FriendlyName = "FriendlyName",
        IconImageName = "Iconimagename",
        RawIcon = BinaryData.FromObjectAsJson("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
        RawPng = BinaryData.FromObjectAsJson("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
        }},
        CertificateName = "certName",
        CertificateExpireOn = DateTimeOffset.Parse("2023-01-02T17:18:19.1234567Z"),
    },
    HostPoolReferences = { },
    KeyVaultUri = new Uri(""),
    FailHealthCheckOnStagingFailure = FailHealthCheckOnStagingFailure.NeedsAssistance,
});
ArmOperation<AppAttachPackageResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, appAttachPackageName, data);
AppAttachPackageResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppAttachPackageData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
