
import com.azure.resourcemanager.mobilenetwork.models.SimGroupResourceId;
import com.azure.resourcemanager.mobilenetwork.models.SimMove;
import java.util.Arrays;

/**
 * Samples for Sims Move.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/SimMove.json
     */
    /**
     * Sample code: Move list of SIMs to target SIM group.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void
        moveListOfSIMsToTargetSIMGroup(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.sims().move("testResourceGroupName", "testSimGroup",
            new SimMove().withTargetSimGroupId(new SimGroupResourceId().withId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg2/providers/Microsoft.MobileNetwork/simGroups/testSimGroup1"))
                .withSims(Arrays.asList("testSim", "testSim2")),
            com.azure.core.util.Context.NONE);
    }
}
