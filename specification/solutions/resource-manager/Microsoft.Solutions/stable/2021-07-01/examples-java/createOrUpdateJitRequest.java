
import com.azure.resourcemanager.managedapplications.models.JitAuthorizationPolicies;
import com.azure.resourcemanager.managedapplications.models.JitSchedulingPolicy;
import com.azure.resourcemanager.managedapplications.models.JitSchedulingType;
import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for JitRequests CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/createOrUpdateJitRequest.
     * json
     */
    /**
     * Sample code: Create or update jit request.
     * 
     * @param manager Entry point to ApplicationManager.
     */
    public static void
        createOrUpdateJitRequest(com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager.jitRequests().define("myJitRequest").withRegion((String) null).withExistingResourceGroup("rg")
            .withApplicationResourceId(
                "/subscriptions/00c76877-e316-48a7-af60-4a09fec9d43f/resourceGroups/52F30DB2/providers/Microsoft.Solutions/applications/7E193158")
            .withJitAuthorizationPolicies(
                Arrays.asList(new JitAuthorizationPolicies().withPrincipalId("1db8e132e2934dbcb8e1178a61319491")
                    .withRoleDefinitionId("ecd05a23-931a-4c38-a52b-ac7c4c583334")))
            .withJitSchedulingPolicy(
                new JitSchedulingPolicy().withType(JitSchedulingType.ONCE).withDuration(Duration.parse("PT8H"))
                    .withStartTime(OffsetDateTime.parse("2021-04-22T05:48:30.6661804Z")))
            .create();
    }
}
