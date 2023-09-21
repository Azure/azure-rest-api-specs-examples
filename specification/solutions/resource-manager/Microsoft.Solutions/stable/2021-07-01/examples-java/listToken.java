import com.azure.resourcemanager.managedapplications.models.ListTokenRequest;
import java.util.Arrays;

/** Samples for Applications ListTokens. */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/listToken.json
     */
    /**
     * Sample code: List tokens for application.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void listTokensForApplication(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager
            .applications()
            .listTokensWithResponse(
                "rg",
                "myManagedApplication",
                new ListTokenRequest()
                    .withAuthorizationAudience("fakeTokenPlaceholder")
                    .withUserAssignedIdentities(Arrays.asList("IdentityOne", "IdentityTwo")),
                com.azure.core.util.Context.NONE);
    }
}
