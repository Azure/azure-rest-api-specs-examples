using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridConnectivity.Models;
using Azure.ResourceManager.HybridConnectivity;

// Generated from example definition: 2024-12-01/SolutionConfigurations_Get.json
// this example is just showing the usage of "SolutionConfiguration_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PublicCloudConnectorSolutionConfigurationResource created on azure
// for more information of creating PublicCloudConnectorSolutionConfigurationResource, please refer to the document of PublicCloudConnectorSolutionConfigurationResource
string resourceUri = "ymuj";
string solutionConfiguration = "tks";
ResourceIdentifier publicCloudConnectorSolutionConfigurationResourceId = PublicCloudConnectorSolutionConfigurationResource.CreateResourceIdentifier(resourceUri, solutionConfiguration);
PublicCloudConnectorSolutionConfigurationResource publicCloudConnectorSolutionConfiguration = client.GetPublicCloudConnectorSolutionConfigurationResource(publicCloudConnectorSolutionConfigurationResourceId);

// invoke the operation
PublicCloudConnectorSolutionConfigurationResource result = await publicCloudConnectorSolutionConfiguration.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PublicCloudConnectorSolutionConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
