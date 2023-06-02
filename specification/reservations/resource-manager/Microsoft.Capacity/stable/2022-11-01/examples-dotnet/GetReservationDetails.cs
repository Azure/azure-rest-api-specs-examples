using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Reservations;
using Azure.ResourceManager.Reservations.Models;

// Generated from example definition: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetReservationDetails.json
// this example is just showing the usage of "Reservation_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ReservationDetailResource created on azure
// for more information of creating ReservationDetailResource, please refer to the document of ReservationDetailResource
Guid reservationOrderId = Guid.Parse("276e7ae4-84d0-4da6-ab4b-d6b94f3557da");
Guid reservationId = Guid.Parse("6ef59113-3482-40da-8d79-787f823e34bc");
ResourceIdentifier reservationDetailResourceId = ReservationDetailResource.CreateResourceIdentifier(reservationOrderId, reservationId);
ReservationDetailResource reservationDetail = client.GetReservationDetailResource(reservationDetailResourceId);

// invoke the operation
string expand = "renewProperties";
ReservationDetailResource result = await reservationDetail.GetAsync(expand: expand);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ReservationDetailData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
