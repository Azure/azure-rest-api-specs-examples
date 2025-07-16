
import com.azure.resourcemanager.playwright.models.CheckNameAvailabilityRequest;

/**
 * Samples for PlaywrightWorkspaces CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01-preview/PlaywrightWorkspaces_CheckNameAvailability.json
     */
    /**
     * Sample code: PlaywrightWorkspaces_CheckNameAvailability.
     * 
     * @param manager Entry point to PlaywrightManager.
     */
    public static void
        playwrightWorkspacesCheckNameAvailability(com.azure.resourcemanager.playwright.PlaywrightManager manager) {
        manager.playwrightWorkspaces().checkNameAvailabilityWithResponse(new CheckNameAvailabilityRequest()
            .withName("dummyName").withType("Microsoft.LoadTestService/PlaywrightWorkspaces"),
            com.azure.core.util.Context.NONE);
    }
}
