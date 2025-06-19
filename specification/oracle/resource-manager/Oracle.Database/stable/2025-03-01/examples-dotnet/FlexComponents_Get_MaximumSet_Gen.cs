using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OracleDatabase;

// Generated from example definition: 2025-03-01/FlexComponents_Get_MaximumSet_Gen.json
// this example is just showing the usage of "FlexComponent_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OracleFlexComponentResource created on azure
// for more information of creating OracleFlexComponentResource, please refer to the document of OracleFlexComponentResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
AzureLocation location = new AzureLocation("eastus");
string flexComponentName = "flexComponent";
ResourceIdentifier oracleFlexComponentResourceId = OracleFlexComponentResource.CreateResourceIdentifier(subscriptionId, location, flexComponentName);
OracleFlexComponentResource oracleFlexComponent = client.GetOracleFlexComponentResource(oracleFlexComponentResourceId);

// invoke the operation
OracleFlexComponentResource result = await oracleFlexComponent.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
OracleFlexComponentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
