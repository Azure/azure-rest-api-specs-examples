
import com.azure.resourcemanager.customerinsights.models.PermissionTypes;
import java.util.Arrays;

/**
 * Samples for AuthorizationPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/
     * AuthorizationPoliciesCreateOrUpdate.json
     */
    /**
     * Sample code: AuthorizationPolicies_CreateOrUpdate.
     * 
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void authorizationPoliciesCreateOrUpdate(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.authorizationPolicies().define("testPolicy4222").withExistingHub("TestHubRG", "azSdkTestHub")
            .withPermissions(Arrays.asList(PermissionTypes.READ, PermissionTypes.WRITE, PermissionTypes.MANAGE))
            .create();
    }
}
