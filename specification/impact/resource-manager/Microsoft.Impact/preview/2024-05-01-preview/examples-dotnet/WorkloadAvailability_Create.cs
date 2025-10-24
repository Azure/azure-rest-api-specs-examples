using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ImpactReporting.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ImpactReporting;

// Generated from example definition: 2024-05-01-preview/WorkloadAvailability_Create.json
// this example is just showing the usage of "WorkloadImpact_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this WorkloadImpactResource
WorkloadImpactCollection collection = subscriptionResource.GetWorkloadImpacts();

// invoke the operation
string workloadImpactName = "impact-002";
WorkloadImpactData data = new WorkloadImpactData
{
    Properties = new WorkloadImpactProperties(DateTimeOffset.Parse("2022-06-15T05:59:46.6517821Z"), new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-rg/providers/Microsoft.Sql/sqlserver/dbservercontext"), "Availability")
    {
        ImpactDescription = "read calls failed",
        Workload = new ImpactedWorkload
        {
            Context = "webapp/scenario1",
            Toolset = ImpactToolset.Other,
        },
        ClientIncidentDetails = new ImpactClientIncidentDetails
        {
            ClientIncidentId = "AA123",
            ClientIncidentSource = ImpactIncidentSource.Jira,
        },
    },
};
ArmOperation<WorkloadImpactResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, workloadImpactName, data);
WorkloadImpactResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WorkloadImpactData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
