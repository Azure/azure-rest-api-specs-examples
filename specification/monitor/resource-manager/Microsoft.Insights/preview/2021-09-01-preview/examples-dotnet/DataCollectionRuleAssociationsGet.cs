using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-09-01-preview/examples/DataCollectionRuleAssociationsGet.json
// this example is just showing the usage of "DataCollectionRuleAssociations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this DataCollectionRuleAssociationResource
string resourceUri = "subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVm";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", resourceUri));
DataCollectionRuleAssociationCollection collection = client.GetDataCollectionRuleAssociations(scopeId);

// invoke the operation
string associationName = "myAssociation";
bool result = await collection.ExistsAsync(associationName);

Console.WriteLine($"Succeeded: {result}");
