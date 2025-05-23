using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Hci.Models;
using Azure.ResourceManager.Hci;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/PutUpdates.json
// this example is just showing the usage of "Updates_Put" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HciClusterResource created on azure
// for more information of creating HciClusterResource, please refer to the document of HciClusterResource
string subscriptionId = "b8d594e5-51f3-4c11-9c54-a7771b81c712";
string resourceGroupName = "testrg";
string clusterName = "testcluster";
ResourceIdentifier hciClusterResourceId = HciClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
HciClusterResource hciCluster = client.GetHciClusterResource(hciClusterResourceId);

// get the collection of this HciClusterUpdateResource
HciClusterUpdateCollection collection = hciCluster.GetHciClusterUpdates();

// invoke the operation
string updateName = "Microsoft4.2203.2.32";
HciClusterUpdateData data = new HciClusterUpdateData
{
    InstalledOn = DateTimeOffset.Parse("2022-04-06T14:08:18.254Z"),
    Description = "AzS Update 4.2203.2.32",
    State = HciUpdateState.Installed,
    Prerequisites = {new HciClusterUpdatePrerequisite
    {
    UpdateType = "update type",
    Version = "prerequisite version",
    PackageName = "update package name",
    }},
    PackagePath = "\\\\SU1FileServer\\SU1_Infrastructure_2\\Updates\\Packages\\Microsoft4.2203.2.32",
    PackageSizeInMb = 18858,
    DisplayName = "AzS Update - 4.2203.2.32",
    Version = "4.2203.2.32",
    Publisher = "Microsoft",
    ReleaseLink = "https://docs.microsoft.com/azure-stack/operator/release-notes?view=azs-2203",
    AvailabilityType = HciAvailabilityType.Local,
    PackageType = "Infrastructure",
    AdditionalProperties = "additional properties",
    ProgressPercentage = 0,
    NotifyMessage = "Brief message with instructions for updates of AvailabilityType Notify",
};
ArmOperation<HciClusterUpdateResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, updateName, data);
HciClusterUpdateResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HciClusterUpdateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
