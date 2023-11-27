using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Blueprint;
using Azure.ResourceManager.Blueprint.Models;

// Generated from example definition: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/ARMTemplateArtifact_Create.json
// this example is just showing the usage of "Artifacts_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BlueprintResource created on azure
// for more information of creating BlueprintResource, please refer to the document of BlueprintResource
string resourceScope = "providers/Microsoft.Management/managementGroups/ContosoOnlineGroup";
string blueprintName = "simpleBlueprint";
ResourceIdentifier blueprintResourceId = BlueprintResource.CreateResourceIdentifier(resourceScope, blueprintName);
BlueprintResource blueprint = client.GetBlueprintResource(blueprintResourceId);

// get the collection of this BlueprintArtifactResource
BlueprintArtifactCollection collection = blueprint.GetBlueprintArtifacts();

// invoke the operation
string artifactName = "storageTemplate";
ArtifactData data = new TemplateArtifact(BinaryData.FromObjectAsJson(new Dictionary<string, object>()
{
    ["contentVersion"] = "1.0.0.0",
    ["outputs"] = new Dictionary<string, object>()
    {
        ["storageAccountName"] = new Dictionary<string, object>()
        {
            ["type"] = "string",
            ["value"] = "[variables('storageAccountName')]"
        }
    },
    ["parameters"] = new Dictionary<string, object>()
    {
        ["storageAccountType"] = new Dictionary<string, object>()
        {
            ["type"] = "string",
            ["allowedValues"] = new object[] { "Standard_LRS", "Standard_GRS", "Standard_ZRS", "Premium_LRS" },
            ["defaultValue"] = "Standard_LRS",
            ["metadata"] = new Dictionary<string, object>()
            {
                ["description"] = "Storage Account type"
            }
        }
    },
    ["resources"] = new object[] { new Dictionary<string, object>()
    {
    ["name"] = "[variables('storageAccountName')]",
    ["type"] = "Microsoft.Storage/storageAccounts",
    ["apiVersion"] = "2016-01-01",
    ["kind"] = "Storage",
    ["location"] = "[resourceGroup().location]",
    ["properties"] = new Dictionary<string, object>()
    {
    },
    ["sku"] = new Dictionary<string, object>()
    {
    ["name"] = "[parameters('storageAccountType')]"}} },
    ["variables"] = new Dictionary<string, object>()
    {
        ["storageAccountName"] = "[concat(uniquestring(resourceGroup().id), 'standardsa')]"
    }
}), new Dictionary<string, ParameterValue>()
{
    ["storageAccountType"] = new ParameterValue()
    {
        Value = BinaryData.FromString("\"[parameters('storageAccountType')]\""),
    },
})
{
    ResourceGroup = "storageRG",
};
ArmOperation<BlueprintArtifactResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, artifactName, data);
BlueprintArtifactResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ArtifactData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
