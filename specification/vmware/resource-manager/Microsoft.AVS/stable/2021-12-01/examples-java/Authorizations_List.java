import com.azure.core.util.Context;

/** Samples for Authorizations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/Authorizations_List.json
     */
    /**
     * Sample code: Authorizations_List.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void authorizationsList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.authorizations().list("group1", "cloud1", Context.NONE);
    }
}
