using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Resources/stable/2022-09-01/examples/ExportResourceGroupWithFiltering.json
// this example is just showing the usage of "ResourceGroups_ExportTemplate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "my-resource-group";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroup = client.GetResourceGroupResource(resourceGroupResourceId);

// invoke the operation
ExportTemplate exportTemplate = new ExportTemplate()
{
    Resources =
    {
    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/my-resource-group/providers/My.RP/myResourceType/myFirstResource"
    },
    Options = "SkipResourceNameParameterization",
};
ArmOperation<ResourceGroupExportResult> lro = await resourceGroup.ExportTemplateAsync(WaitUntil.Completed, exportTemplate);
ResourceGroupExportResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
