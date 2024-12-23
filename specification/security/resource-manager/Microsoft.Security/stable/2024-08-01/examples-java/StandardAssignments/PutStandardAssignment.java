
import com.azure.resourcemanager.security.models.AssignedStandardItem;
import com.azure.resourcemanager.security.models.Effect;
import java.util.Arrays;

/**
 * Samples for StandardAssignments Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-08-01/examples/StandardAssignments/
     * PutStandardAssignment.json
     */
    /**
     * Sample code: Put an audit standard assignment.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void putAnAuditStandardAssignment(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.standardAssignments().define("1f3afdf9-d0c9-4c3d-847f-89da613e70a8")
            .withExistingResourceId("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23").withDisplayName("ASC Default")
            .withDescription("Set of policies monitored by Azure Security Center for cross cloud")
            .withAssignedStandard(new AssignedStandardItem()
                .withId("/providers/Microsoft.Security/securityStandards/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"))
            .withEffect(Effect.AUDIT).withExcludedScopes(Arrays.asList()).create();
    }
}
