using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/ManagedNetwork/getRule.json
// this example is just showing the usage of "ManagedNetworkSettingsRule_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningOutboundRuleBasicResource created on azure
// for more information of creating MachineLearningOutboundRuleBasicResource, please refer to the document of MachineLearningOutboundRuleBasicResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "aml-workspace-name";
string ruleName = "some_string";
ResourceIdentifier machineLearningOutboundRuleBasicResourceId = MachineLearningOutboundRuleBasicResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, ruleName);
MachineLearningOutboundRuleBasicResource machineLearningOutboundRuleBasic = client.GetMachineLearningOutboundRuleBasicResource(machineLearningOutboundRuleBasicResourceId);

// invoke the operation
MachineLearningOutboundRuleBasicResource result = await machineLearningOutboundRuleBasic.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningOutboundRuleBasicData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
