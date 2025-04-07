using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AgriculturePlatform.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.AgriculturePlatform;

// Generated from example definition: 2024-06-01-preview/AgriService_Update_MaximumSet_Gen.json
// this example is just showing the usage of "AgriServiceResource_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AgricultureServiceResource created on azure
// for more information of creating AgricultureServiceResource, please refer to the document of AgricultureServiceResource
string subscriptionId = "83D293F5-DEFD-4D48-B120-1DC713BE338A";
string resourceGroupName = "rgopenapi";
string agriServiceResourceName = "abc123";
ResourceIdentifier agricultureServiceResourceId = AgricultureServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, agriServiceResourceName);
AgricultureServiceResource agricultureService = client.GetAgricultureServiceResource(agricultureServiceResourceId);

// invoke the operation
AgricultureServicePatch patch = new AgricultureServicePatch
{
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("key4771")] = new UserAssignedIdentity()
        },
    },
    Sku = new AgriculturePlatformSku("tbdtdfffkar")
    {
        Tier = AgriculturePlatformSkuTier.Free,
        Size = "iusaqqj",
        Family = "hxojswlgs",
        Capacity = 22,
    },
    Tags =
    {
    ["key9006"] = "kuzlwpujbql"
    },
    Properties = new AgricultureServicePatchProperties
    {
        Config = new AgricultureServiceConfig(),
        DataConnectorCredentials = {new DataConnectorCredentialMap("BackendAADApplicationCredentials", new DataConnectorCredentials
        {
        ClientId = "dce298a8-1eec-481a-a8f9-a3cd5a8257b2",
        })},
        InstalledSolutions = {new InstalledSolutionMap("bayerAgPowered.cwum", new AgricultureSolution
        {
        ApplicationName = "bayerAgPowered.cwum",
        })},
    },
};
ArmOperation<AgricultureServiceResource> lro = await agricultureService.UpdateAsync(WaitUntil.Completed, patch);
AgricultureServiceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AgricultureServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
