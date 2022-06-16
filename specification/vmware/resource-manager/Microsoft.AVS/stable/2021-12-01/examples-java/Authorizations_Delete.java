import com.azure.core.util.Context;

/** Samples for Authorizations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/Authorizations_Delete.json
     */
    /**
     * Sample code: Authorizations_Delete.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void authorizationsDelete(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.authorizations().delete("group1", "cloud1", "authorization1", Context.NONE);
    }
}
