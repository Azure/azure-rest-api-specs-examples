using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/GetAseCustomDnsSuffixConfiguration.json
// this example is just showing the usage of "AppServiceEnvironments_GetAseCustomDnsSuffixConfiguration" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CustomDnsSuffixConfigurationResource created on azure
// for more information of creating CustomDnsSuffixConfigurationResource, please refer to the document of CustomDnsSuffixConfigurationResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "test-rg";
string name = "test-ase";
ResourceIdentifier customDnsSuffixConfigurationResourceId = CustomDnsSuffixConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
CustomDnsSuffixConfigurationResource customDnsSuffixConfiguration = client.GetCustomDnsSuffixConfigurationResource(customDnsSuffixConfigurationResourceId);

// invoke the operation
CustomDnsSuffixConfigurationResource result = await customDnsSuffixConfiguration.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CustomDnsSuffixConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
