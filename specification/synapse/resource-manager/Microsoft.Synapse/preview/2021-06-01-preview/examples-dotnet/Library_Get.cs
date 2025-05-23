using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/Library_Get.json
// this example is just showing the usage of "Library_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseLibraryResource created on azure
// for more information of creating SynapseLibraryResource, please refer to the document of SynapseLibraryResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string workspaceName = "exampleWorkspace";
string libraryName = "exampleLibraryName.jar";
ResourceIdentifier synapseLibraryResourceId = SynapseLibraryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, libraryName);
SynapseLibraryResource synapseLibrary = client.GetSynapseLibraryResource(synapseLibraryResourceId);

// invoke the operation
SynapseLibraryResource result = await synapseLibrary.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SynapseLibraryData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
