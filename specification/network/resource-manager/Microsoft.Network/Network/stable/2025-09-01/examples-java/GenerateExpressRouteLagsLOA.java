
import com.azure.resourcemanager.network.models.GenerateExpressRouteLagsLOARequest;
import java.util.Arrays;

/**
 * Samples for ExpressRouteLags GenerateLoa.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/GenerateExpressRouteLagsLOA.json
     */
    /**
     * Sample code: Generate express route lag LOA.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void generateExpressRouteLagLOA(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteLags()
            .generateLoaWithResponse(
                "rg1", "lagName", new GenerateExpressRouteLagsLOARequest().withCustomerName("Customer Name")
                    .withMembers(Arrays.asList("member1", "member2", "member3", "member4")),
                com.azure.core.util.Context.NONE);
    }
}
