import com.azure.core.util.Context;
import com.azure.resourcemanager.mobilenetwork.models.SimDeleteList;
import java.util.Arrays;

/** Samples for SimOperation BulkDelete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/SimBulkDelete.json
     */
    /**
     * Sample code: Bulk delete SIMs from a SIM group.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void bulkDeleteSIMsFromASIMGroup(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .simOperations()
            .bulkDelete(
                "testResourceGroupName",
                "testSimGroup",
                new SimDeleteList().withSims(Arrays.asList("testSim", "testSim2")),
                Context.NONE);
    }
}
