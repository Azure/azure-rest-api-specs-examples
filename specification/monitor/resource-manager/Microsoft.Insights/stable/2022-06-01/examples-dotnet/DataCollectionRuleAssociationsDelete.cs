using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/DataCollectionRuleAssociationsDelete.json
// this example is just showing the usage of "DataCollectionRuleAssociations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataCollectionRuleAssociationResource created on azure
// for more information of creating DataCollectionRuleAssociationResource, please refer to the document of DataCollectionRuleAssociationResource
string resourceUri = "subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVm";
string associationName = "myAssociation";
ResourceIdentifier dataCollectionRuleAssociationResourceId = DataCollectionRuleAssociationResource.CreateResourceIdentifier(resourceUri, associationName);
DataCollectionRuleAssociationResource dataCollectionRuleAssociation = client.GetDataCollectionRuleAssociationResource(dataCollectionRuleAssociationResourceId);

// invoke the operation
await dataCollectionRuleAssociation.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
