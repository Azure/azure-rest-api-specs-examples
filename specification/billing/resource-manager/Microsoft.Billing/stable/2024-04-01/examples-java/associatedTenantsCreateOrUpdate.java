
import com.azure.resourcemanager.billing.fluent.models.AssociatedTenantInner;
import com.azure.resourcemanager.billing.models.AssociatedTenantProperties;
import com.azure.resourcemanager.billing.models.BillingManagementTenantState;
import com.azure.resourcemanager.billing.models.ProvisioningTenantState;

/**
 * Samples for AssociatedTenants CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * associatedTenantsCreateOrUpdate.json
     */
    /**
     * Sample code: AssociatedTenantsCreateOrUpdate.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void associatedTenantsCreateOrUpdate(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.associatedTenants().createOrUpdate(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            "11111111-1111-1111-1111-111111111111",
            new AssociatedTenantInner().withProperties(new AssociatedTenantProperties()
                .withDisplayName("Contoso Finance").withBillingManagementState(BillingManagementTenantState.ACTIVE)
                .withProvisioningManagementState(ProvisioningTenantState.PENDING)),
            com.azure.core.util.Context.NONE);
    }
}
