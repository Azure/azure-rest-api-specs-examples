using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Reservations;
using Azure.ResourceManager.Reservations.Models;

// Generated from example definition: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/GetAvailableScope.json
// this example is just showing the usage of "Reservation_AvailableScopes" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ReservationDetailResource created on azure
// for more information of creating ReservationDetailResource, please refer to the document of ReservationDetailResource
Guid reservationOrderId = Guid.Parse("276e7ae4-84d0-4da6-ab4b-d6b94f3557da");
Guid reservationId = Guid.Parse("356e7ae4-84d0-4da6-ab4b-d6b94f3557da");
ResourceIdentifier reservationDetailResourceId = ReservationDetailResource.CreateResourceIdentifier(reservationOrderId, reservationId);
ReservationDetailResource reservationDetail = client.GetReservationDetailResource(reservationDetailResourceId);

// invoke the operation
AvailableScopesContent content = new AvailableScopesContent()
{
    Scopes =
    {
    "/subscriptions/efc7c997-7700-4a74-b731-55aec16c15e9"
    },
};
ArmOperation<AvailableScopesProperties> lro = await reservationDetail.GetAvailableScopesAsync(WaitUntil.Completed, content);
AvailableScopesProperties result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
