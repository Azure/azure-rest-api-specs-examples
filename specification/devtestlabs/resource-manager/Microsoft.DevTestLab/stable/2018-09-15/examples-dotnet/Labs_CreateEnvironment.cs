using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevTestLabs;
using Azure.ResourceManager.DevTestLabs.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_CreateEnvironment.json
// this example is just showing the usage of "Labs_CreateEnvironment" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabResource created on azure
// for more information of creating DevTestLabResource, please refer to the document of DevTestLabResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string name = "{labName}";
ResourceIdentifier devTestLabResourceId = DevTestLabResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
DevTestLabResource devTestLab = client.GetDevTestLabResource(devTestLabResourceId);

// invoke the operation
DevTestLabVmCreationContent content = new DevTestLabVmCreationContent()
{
    Name = "{vmName}",
    Location = new AzureLocation("{location}"),
    Tags =
    {
    ["tagName1"] = "tagValue1",
    },
    Size = "Standard_A2_v2",
    UserName = "{userName}",
    Password = "{userPassword}",
    LabSubnetName = "{virtualnetwork-subnet-name}",
    LabVirtualNetworkId = new ResourceIdentifier("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualnetworks/{virtualNetworkName}"),
    DisallowPublicIPAddress = true,
    GalleryImageReference = new DevTestLabGalleryImageReference()
    {
        Offer = "UbuntuServer",
        Publisher = "Canonical",
        Sku = "16.04-LTS",
        OSType = "Linux",
        Version = "Latest",
    },
    AllowClaim = true,
    StorageType = "Standard",
};
await devTestLab.CreateEnvironmentAsync(WaitUntil.Completed, content);

Console.WriteLine($"Succeeded");
