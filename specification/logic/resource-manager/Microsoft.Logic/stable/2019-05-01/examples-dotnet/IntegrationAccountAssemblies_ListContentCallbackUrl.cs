using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Logic.Models;
using Azure.ResourceManager.Logic;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountAssemblies_ListContentCallbackUrl.json
// this example is just showing the usage of "IntegrationAccountAssemblies_ListContentCallbackUrl" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IntegrationAccountAssemblyDefinitionResource created on azure
// for more information of creating IntegrationAccountAssemblyDefinitionResource, please refer to the document of IntegrationAccountAssemblyDefinitionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testResourceGroup";
string integrationAccountName = "testIntegrationAccount";
string assemblyArtifactName = "testAssembly";
ResourceIdentifier integrationAccountAssemblyDefinitionResourceId = IntegrationAccountAssemblyDefinitionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, integrationAccountName, assemblyArtifactName);
IntegrationAccountAssemblyDefinitionResource integrationAccountAssemblyDefinition = client.GetIntegrationAccountAssemblyDefinitionResource(integrationAccountAssemblyDefinitionResourceId);

// invoke the operation
LogicWorkflowTriggerCallbackUri result = await integrationAccountAssemblyDefinition.GetContentCallbackUrlAsync();

Console.WriteLine($"Succeeded: {result}");
