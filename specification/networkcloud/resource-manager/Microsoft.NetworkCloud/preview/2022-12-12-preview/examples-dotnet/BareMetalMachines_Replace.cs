using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BareMetalMachines_Replace.json
// this example is just showing the usage of "BareMetalMachines_Replace" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BareMetalMachineResource created on azure
// for more information of creating BareMetalMachineResource, please refer to the document of BareMetalMachineResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string bareMetalMachineName = "bareMetalMachineName";
ResourceIdentifier bareMetalMachineResourceId = BareMetalMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, bareMetalMachineName);
BareMetalMachineResource bareMetalMachine = client.GetBareMetalMachineResource(bareMetalMachineResourceId);

// invoke the operation
BareMetalMachineReplaceContent content = new BareMetalMachineReplaceContent()
{
    BmcCredentials = new AdministrativeCredentials("bmcuser")
    {
        Password = "{password}",
    },
    BmcMacAddress = "00:00:4f:00:57:ad",
    BootMacAddress = "00:00:4e:00:58:af",
    MachineName = "name",
    SerialNumber = "BM1219XXX",
};
await bareMetalMachine.ReplaceAsync(WaitUntil.Completed, content: content);

Console.WriteLine($"Succeeded");
