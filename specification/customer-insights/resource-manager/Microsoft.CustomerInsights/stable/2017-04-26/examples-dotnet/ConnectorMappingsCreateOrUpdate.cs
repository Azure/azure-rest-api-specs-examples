using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CustomerInsights.Models;
using Azure.ResourceManager.CustomerInsights;

// Generated from example definition: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ConnectorMappingsCreateOrUpdate.json
// this example is just showing the usage of "ConnectorMappings_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ConnectorMappingResourceFormatResource created on azure
// for more information of creating ConnectorMappingResourceFormatResource, please refer to the document of ConnectorMappingResourceFormatResource
string subscriptionId = "subid";
string resourceGroupName = "TestHubRG";
string hubName = "sdkTestHub";
string connectorName = "testConnector8858";
string mappingName = "testMapping12491";
ResourceIdentifier connectorMappingResourceFormatResourceId = ConnectorMappingResourceFormatResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hubName, connectorName, mappingName);
ConnectorMappingResourceFormatResource connectorMappingResourceFormat = client.GetConnectorMappingResourceFormatResource(connectorMappingResourceFormatResourceId);

// invoke the operation
ConnectorMappingResourceFormatData data = new ConnectorMappingResourceFormatData
{
    EntityType = EntityType.Interaction,
    EntityTypeName = "TestInteractionType2967",
    DisplayName = "testMapping12491",
    Description = "Test mapping",
    MappingProperties = new ConnectorMappingProperties(new ConnectorMappingErrorManagement(ErrorManagementType.StopImport)
    {
        ErrorLimit = 10,
    }, new ConnectorMappingFormat
    {
        ColumnDelimiter = "|",
    }, new ConnectorMappingAvailability(5)
    {
        Frequency = FrequencyType.Hour,
    }, new ConnectorMappingStructure[]
{
new ConnectorMappingStructure("unknwon1", "unknown1")
{
IsEncrypted = false,
},
new ConnectorMappingStructure("unknwon2", "unknown2")
{
IsEncrypted = true,
}
}, new ConnectorMappingCompleteOperation
{
    CompletionOperationType = CompletionOperationType.DeleteFile,
    DestinationFolder = "fakePath",
})
    {
        FolderPath = "http://sample.dne/file",
        FileFilter = "unknown",
        HasHeader = false,
    },
};
ArmOperation<ConnectorMappingResourceFormatResource> lro = await connectorMappingResourceFormat.UpdateAsync(WaitUntil.Completed, data);
ConnectorMappingResourceFormatResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ConnectorMappingResourceFormatData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
