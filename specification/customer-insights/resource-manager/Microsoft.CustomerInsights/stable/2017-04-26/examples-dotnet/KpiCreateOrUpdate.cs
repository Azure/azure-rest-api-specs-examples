using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CustomerInsights;
using Azure.ResourceManager.CustomerInsights.Models;

// Generated from example definition: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/KpiCreateOrUpdate.json
// this example is just showing the usage of "Kpi_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HubResource created on azure
// for more information of creating HubResource, please refer to the document of HubResource
string subscriptionId = "subid";
string resourceGroupName = "TestHubRG";
string hubName = "sdkTestHub";
ResourceIdentifier hubResourceId = HubResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hubName);
HubResource hub = client.GetHubResource(hubResourceId);

// get the collection of this KpiResourceFormatResource
KpiResourceFormatCollection collection = hub.GetKpiResourceFormats();

// invoke the operation
string kpiName = "kpiTest45453647";
KpiResourceFormatData data = new KpiResourceFormatData()
{
    EntityType = EntityType.Profile,
    EntityTypeName = "testProfile2327128",
    DisplayName =
    {
    ["en-us"] = "Kpi DisplayName",
    },
    Description =
    {
    ["en-us"] = "Kpi Description",
    },
    CalculationWindow = CalculationWindowType.Day,
    Function = KpiFunction.Sum,
    Expression = "SavingAccountBalance",
    Unit = "unit",
    GroupBy =
    {
    "SavingAccountBalance"
    },
    ThresHolds = new KpiThresholds(5, 50, true),
    Aliases =
    {
    new KpiAlias("alias","Id+4")
    },
};
ArmOperation<KpiResourceFormatResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, kpiName, data);
KpiResourceFormatResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KpiResourceFormatData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
