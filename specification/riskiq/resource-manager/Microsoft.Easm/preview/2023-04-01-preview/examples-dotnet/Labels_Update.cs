using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DefenderEasm.Models;
using Azure.ResourceManager.DefenderEasm;

// Generated from example definition: specification/riskiq/resource-manager/Microsoft.Easm/preview/2023-04-01-preview/examples/Labels_Update.json
// this example is just showing the usage of "Labels_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EasmLabelResource created on azure
// for more information of creating EasmLabelResource, please refer to the document of EasmLabelResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "dummyrg";
string workspaceName = "ThisisaWorkspace";
string labelName = "ThisisaLabel";
ResourceIdentifier easmLabelResourceId = EasmLabelResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, labelName);
EasmLabelResource easmLabel = client.GetEasmLabelResource(easmLabelResourceId);

// invoke the operation
EasmLabelPatch patch = new EasmLabelPatch();
EasmLabelResource result = await easmLabel.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EasmLabelData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
