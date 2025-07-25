using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/ProjectConnection/update.json
// this example is just showing the usage of "ProjectConnections_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CognitiveServicesProjectConnectionResource created on azure
// for more information of creating CognitiveServicesProjectConnectionResource, please refer to the document of CognitiveServicesProjectConnectionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string accountName = "account-1";
string projectName = "project-1";
string connectionName = "connection-1";
ResourceIdentifier cognitiveServicesProjectConnectionResourceId = CognitiveServicesProjectConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, projectName, connectionName);
CognitiveServicesProjectConnectionResource cognitiveServicesProjectConnection = client.GetCognitiveServicesProjectConnectionResource(cognitiveServicesProjectConnectionResourceId);

// invoke the operation
CognitiveServicesConnectionPatch patch = new CognitiveServicesConnectionPatch
{
    Properties = new AccessKeyAuthTypeConnectionProperties
    {
        Credentials = new CognitiveServicesConnectionAccessKey
        {
            AccessKeyId = "some_string",
            SecretAccessKey = "some_string",
        },
        Category = CognitiveServicesConnectionCategory.AdlsGen2,
        ExpiryOn = DateTimeOffset.Parse("2020-01-01T00:00:00Z"),
        Metadata = { },
        Target = "some_string",
    },
};
CognitiveServicesProjectConnectionResource result = await cognitiveServicesProjectConnection.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CognitiveServicesConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
