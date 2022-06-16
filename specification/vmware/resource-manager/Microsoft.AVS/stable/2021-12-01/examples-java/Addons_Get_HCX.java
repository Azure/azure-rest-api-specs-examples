import com.azure.core.util.Context;

/** Samples for Addons Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/Addons_Get_HCX.json
     */
    /**
     * Sample code: Addons_Get_HCX.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void addonsGetHCX(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.addons().getWithResponse("group1", "cloud1", "hcx", Context.NONE);
    }
}
