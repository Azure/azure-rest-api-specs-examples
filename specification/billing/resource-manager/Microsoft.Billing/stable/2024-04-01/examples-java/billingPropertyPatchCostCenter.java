
import com.azure.resourcemanager.billing.fluent.models.BillingPropertyInner;
import com.azure.resourcemanager.billing.models.BillingPropertyProperties;

/**
 * Samples for BillingProperty Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingPropertyPatchCostCenter.json
     */
    /**
     * Sample code: BillingPropertyPatchCostCenter.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingPropertyPatchCostCenter(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingProperties().updateWithResponse(
            new BillingPropertyInner().withProperties(new BillingPropertyProperties().withCostCenter("1010")),
            com.azure.core.util.Context.NONE);
    }
}
