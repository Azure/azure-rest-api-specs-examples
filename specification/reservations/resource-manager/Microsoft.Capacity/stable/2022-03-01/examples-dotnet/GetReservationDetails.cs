using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Reservations;

// Generated from example definition: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/GetReservationDetails.json
// this example is just showing the usage of "Reservation_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ReservationOrderResource created on azure
// for more information of creating ReservationOrderResource, please refer to the document of ReservationOrderResource
Guid reservationOrderId = Guid.Parse("276e7ae4-84d0-4da6-ab4b-d6b94f3557da");
ResourceIdentifier reservationOrderResourceId = ReservationOrderResource.CreateResourceIdentifier(reservationOrderId);
ReservationOrderResource reservationOrder = client.GetReservationOrderResource(reservationOrderResourceId);

// get the collection of this ReservationDetailResource
ReservationDetailCollection collection = reservationOrder.GetReservationDetails();

// invoke the operation
Guid reservationId = Guid.Parse("6ef59113-3482-40da-8d79-787f823e34bc");
string expand = "renewProperties";
bool result = await collection.ExistsAsync(reservationId, expand: expand);

Console.WriteLine($"Succeeded: {result}");
