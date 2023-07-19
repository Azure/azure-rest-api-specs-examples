import com.azure.resourcemanager.mobilenetwork.models.SimDeleteList;
import java.util.Arrays;

/** Samples for Sims BulkDelete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/SimBulkDelete.json
     */
    /**
     * Sample code: Bulk delete SIMs from a SIM group.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void bulkDeleteSIMsFromASIMGroup(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .sims()
            .bulkDelete(
                "testResourceGroupName",
                "testSimGroup",
                new SimDeleteList().withSims(Arrays.asList("testSim", "testSim2")),
                com.azure.core.util.Context.NONE);
    }
}
