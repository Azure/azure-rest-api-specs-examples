using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Synapse.Models;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/DeleteBigDataPool.json
// this example is just showing the usage of "BigDataPools_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseBigDataPoolInfoResource created on azure
// for more information of creating SynapseBigDataPoolInfoResource, please refer to the document of SynapseBigDataPoolInfoResource
string subscriptionId = "01234567-89ab-4def-0123-456789abcdef";
string resourceGroupName = "ExampleResourceGroup";
string workspaceName = "ExampleWorkspace";
string bigDataPoolName = "ExamplePool";
ResourceIdentifier synapseBigDataPoolInfoResourceId = SynapseBigDataPoolInfoResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, bigDataPoolName);
SynapseBigDataPoolInfoResource synapseBigDataPoolInfo = client.GetSynapseBigDataPoolInfoResource(synapseBigDataPoolInfoResourceId);

// invoke the operation
ArmOperation<BinaryData> lro = await synapseBigDataPoolInfo.DeleteAsync(WaitUntil.Completed);
BinaryData result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
