using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Datadog;

// Generated from example definition: specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/SingleSignOnConfigurations_Get.json
// this example is just showing the usage of "SingleSignOnConfigurations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DatadogSingleSignOnResource created on azure
// for more information of creating DatadogSingleSignOnResource, please refer to the document of DatadogSingleSignOnResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "myMonitor";
string configurationName = "default";
ResourceIdentifier datadogSingleSignOnResourceId = DatadogSingleSignOnResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName, configurationName);
DatadogSingleSignOnResource datadogSingleSignOnResource = client.GetDatadogSingleSignOnResource(datadogSingleSignOnResourceId);

// invoke the operation
DatadogSingleSignOnResource result = await datadogSingleSignOnResource.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DatadogSingleSignOnResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
