using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Workloads;

// Generated from example definition: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPApplicationServerInstances_Create_HA_AvSet.json
// this example is just showing the usage of "SAPApplicationServerInstances_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SapVirtualInstanceResource created on azure
// for more information of creating SapVirtualInstanceResource, please refer to the document of SapVirtualInstanceResource
string subscriptionId = "6d875e77-e412-4d7d-9af4-8895278b4443";
string resourceGroupName = "test-rg";
string sapVirtualInstanceName = "X00";
ResourceIdentifier sapVirtualInstanceResourceId = SapVirtualInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sapVirtualInstanceName);
SapVirtualInstanceResource sapVirtualInstance = client.GetSapVirtualInstanceResource(sapVirtualInstanceResourceId);

// get the collection of this SapApplicationServerInstanceResource
SapApplicationServerInstanceCollection collection = sapVirtualInstance.GetSapApplicationServerInstances();

// invoke the operation
string applicationInstanceName = "app01";
SapApplicationServerInstanceData data = new SapApplicationServerInstanceData(new AzureLocation("westcentralus"))
{
    Tags =
    {
    },
};
ArmOperation<SapApplicationServerInstanceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, applicationInstanceName, data);
SapApplicationServerInstanceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SapApplicationServerInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
