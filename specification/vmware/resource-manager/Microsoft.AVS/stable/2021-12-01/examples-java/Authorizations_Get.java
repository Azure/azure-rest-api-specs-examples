import com.azure.core.util.Context;

/** Samples for Authorizations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/Authorizations_Get.json
     */
    /**
     * Sample code: Authorizations_Get.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void authorizationsGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.authorizations().getWithResponse("group1", "cloud1", "authorization1", Context.NONE);
    }
}
