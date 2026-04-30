using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PureStorageBlock.Models;
using Azure.ResourceManager.PureStorageBlock;

// Generated from example definition: 2024-11-01/Reservations_Update_MaximumSet_Gen.json
// this example is just showing the usage of "Reservation_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PureStorageReservationResource created on azure
// for more information of creating PureStorageReservationResource, please refer to the document of PureStorageReservationResource
string subscriptionId = "BC47D6CC-AA80-4374-86F8-19D94EC70666";
string resourceGroupName = "rgpurestorage";
string reservationName = "storagePoolname";
ResourceIdentifier pureStorageReservationResourceId = PureStorageReservationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, reservationName);
PureStorageReservationResource pureStorageReservation = client.GetPureStorageReservationResource(pureStorageReservationResourceId);

// invoke the operation
PureStorageReservationPatch patch = new PureStorageReservationPatch
{
    Tags =
    {
    ["key8751"] = "oikntqrti"
    },
    ReservationUpdateUser = new PureStorageUserDetails("sjzquetrvxcrajxdfwfeuro", "qimvqxnlbclfouwzfk", "abc@example.com")
    {
        Upn = "pvafwnbigmhuigxfu",
        PhoneNumber = "jfljnoxsfsplwczwgvmlurfnorimvl",
        CompanyDetails = new PureStorageCompanyDetails("uleytbkckdhaiykwjjcjqmnlik")
        {
            Address = new PureStorageAddressDetails("ryaasdffnhwialrgmukpiwtcjgbb", "kdpzfxfbgozxwunkkhjthqdsnmce", "fygrbnektar", "trmpjpxsfmxprlnv", "yjttfktdxdzhsgomefhcn")
            {
                AddressLine2 = "cvyuuqnvuqfrpkoplfzmhnwrqsbsgn",
            },
        },
    },
};
ArmOperation<PureStorageReservationResource> lro = await pureStorageReservation.UpdateAsync(WaitUntil.Completed, patch);
PureStorageReservationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PureStorageReservationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
