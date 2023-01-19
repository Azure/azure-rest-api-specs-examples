using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Reservations;
using Azure.ResourceManager.Reservations.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/putMachineLearningServicesQuotaRequestDedicated.json
// this example is just showing the usage of "Quota_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "D7EC67B3-7657-4966-BFFC-41EFD36BAAB3";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this ReservationQuotaResource
string providerId = "Microsoft.MachineLearningServices";
AzureLocation location = new AzureLocation("eastus");
ReservationQuotaCollection collection = subscriptionResource.GetAllReservationQuota(providerId, location);

// invoke the operation
string resourceName = "StandardDv2Family";
ReservationQuotaData data = new ReservationQuotaData()
{
    Properties = new QuotaProperties()
    {
        Limit = 200,
        Unit = "Count",
        ResourceName = new ReservationResourceName()
        {
            Value = "StandardDv2Family",
        },
        ResourceTypeName = ResourceTypeName.Dedicated,
    },
};
ArmOperation<ReservationQuotaResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceName, data);
ReservationQuotaResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ReservationQuotaData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
