using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PureStorageBlock.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.PureStorageBlock;

// Generated from example definition: 2024-11-01/Reservations_Create_MaximumSet_Gen.json
// this example is just showing the usage of "Reservation_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "BC47D6CC-AA80-4374-86F8-19D94EC70666";
string resourceGroupName = "rgpurestorage";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this PureStorageReservationResource
PureStorageReservationCollection collection = resourceGroupResource.GetPureStorageReservations();

// invoke the operation
string reservationName = "storagePoolname";
PureStorageReservationData data = new PureStorageReservationData(new AzureLocation("jynnbjysbc"))
{
    Properties = new PureStorageReservationProperties(new PureStorageMarketplaceDetails(new PureStorageOfferDetails("vejockfhoavaqjvhtwvctdnaefvw", "efojrbphbimq", "caj")
    {
        PlanName = "lvvzchm",
        TermUnit = "ose",
        TermId = "ucyvzkedohfjazifxweylhnbcmeza",
    })
    {
        SubscriptionStatus = PureStorageMarketplaceSubscriptionStatus.PendingFulfillmentStart,
    }, new PureStorageUserDetails("bucysqbbclhwxrzig", "fnsvxlop", "abc@example.com")
    {
        Upn = "ekqbqgpdylggddusuiifrnjcwiefay",
        PhoneNumber = "jglihtgsacdxocc",
        CompanyDetails = new PureStorageCompanyDetails("nrndfzmrakk")
        {
            Address = new PureStorageAddressDetails("f", "qxzhxjoatyuajoljfkd", "dnusygshfvmebpmcjsd", "nuexbknolfphlfguyzq", "yjzqichkfffbdtcswzolmrl")
            {
                AddressLine2 = "gycfosmknj",
            },
        },
    }),
    Tags =
    {
    ["key1110"] = "euhfdmtfpucwurtu"
    },
};
ArmOperation<PureStorageReservationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, reservationName, data);
PureStorageReservationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PureStorageReservationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
