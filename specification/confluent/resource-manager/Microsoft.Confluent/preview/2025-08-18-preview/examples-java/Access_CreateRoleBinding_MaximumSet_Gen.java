
import com.azure.resourcemanager.confluent.models.AccessCreateRoleBindingRequestModel;

/**
 * Samples for Access CreateRoleBinding.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Access_CreateRoleBinding_MaximumSet_Gen.json
     */
    /**
     * Sample code: Access_CreateRoleBinding_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void accessCreateRoleBindingMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.access().createRoleBindingWithResponse(
            "rgconfluent", "ablufxskoyvgtbngsfexfkdw", new AccessCreateRoleBindingRequestModel()
                .withPrincipal("xzbkopaxz").withRoleName("dhqxbrapwgqnmpbrredgxa").withCrnPattern("iif"),
            com.azure.core.util.Context.NONE);
    }
}
