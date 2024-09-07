
import com.azure.resourcemanager.billing.models.AcceptTransferRequest;
import com.azure.resourcemanager.billing.models.ProductDetails;
import com.azure.resourcemanager.billing.models.ProductType;
import java.util.Arrays;

/**
 * Samples for RecipientTransfers Validate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/recipientTransfersValidate.
     * json
     */
    /**
     * Sample code: ValidateTransfer.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void validateTransfer(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.recipientTransfers().validateWithResponse("aabb123",
            new AcceptTransferRequest().withProductDetails(Arrays.asList(
                new ProductDetails().withProductType(ProductType.AZURE_SUBSCRIPTION).withProductId("subscriptionId"),
                new ProductDetails().withProductType(ProductType.AZURE_RESERVATION)
                    .withProductId("reservedInstanceId"))),
            com.azure.core.util.Context.NONE);
    }
}
