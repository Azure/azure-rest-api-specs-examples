
import com.azure.resourcemanager.security.models.AlertSimulatorBundlesRequestProperties;
import com.azure.resourcemanager.security.models.AlertSimulatorRequestBody;
import com.azure.resourcemanager.security.models.BundleType;
import java.util.Arrays;

/**
 * Samples for Alerts Simulate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2022-01-01/examples/Alerts/
     * SimulateAlerts_example.json
     */
    /**
     * Sample code: Simulate security alerts on a subscription.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        simulateSecurityAlertsOnASubscription(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.alerts().simulate("centralus",
            new AlertSimulatorRequestBody().withProperties(new AlertSimulatorBundlesRequestProperties()
                .withBundles(Arrays.asList(BundleType.APP_SERVICES, BundleType.DNS, BundleType.KEY_VAULTS,
                    BundleType.KUBERNETES_SERVICE, BundleType.RESOURCE_MANAGER, BundleType.SQL_SERVERS,
                    BundleType.STORAGE_ACCOUNTS, BundleType.VIRTUAL_MACHINES, BundleType.COSMOS_DBS))),
            com.azure.core.util.Context.NONE);
    }
}
