
import com.azure.resourcemanager.confluent.models.AccessCreateRoleBindingRequestModel;

/**
 * Samples for Access CreateRoleBinding.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Access_CreateRoleBinding_MinimumSet_Gen.json
     */
    /**
     * Sample code: Access_CreateRoleBinding_MinimumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void accessCreateRoleBindingMinimumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.access().createRoleBindingWithResponse("rgconfluent", "gdzfl",
            new AccessCreateRoleBindingRequestModel(), com.azure.core.util.Context.NONE);
    }
}
