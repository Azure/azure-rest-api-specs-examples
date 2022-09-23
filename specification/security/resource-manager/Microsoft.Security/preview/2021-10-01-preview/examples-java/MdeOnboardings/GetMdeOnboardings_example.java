import com.azure.core.util.Context;

/** Samples for MdeOnboardings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-10-01-preview/examples/MdeOnboardings/GetMdeOnboardings_example.json
     */
    /**
     * Sample code: The default configuration or data needed to onboard the machine to MDE.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void theDefaultConfigurationOrDataNeededToOnboardTheMachineToMDE(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.mdeOnboardings().getWithResponse(Context.NONE);
    }
}
