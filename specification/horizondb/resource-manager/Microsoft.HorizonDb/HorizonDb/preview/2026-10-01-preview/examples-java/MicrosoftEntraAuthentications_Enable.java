
import com.azure.resourcemanager.horizondb.models.AuthenticationState;
import com.azure.resourcemanager.horizondb.models.MicrosoftEntraAuthenticationForPatchUpdate;
import com.azure.resourcemanager.horizondb.models.MicrosoftEntraAuthenticationPropertiesForPatchUpdate;

/**
 * Samples for MicrosoftEntraAuthentications Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/MicrosoftEntraAuthentications_Enable.json
     */
    /**
     * Sample code: Start enabling Microsoft Entra authentication.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        startEnablingMicrosoftEntraAuthentication(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.microsoftEntraAuthentications().update("example-resource-group", "example-cluster",
            new MicrosoftEntraAuthenticationForPatchUpdate().withProperties(
                new MicrosoftEntraAuthenticationPropertiesForPatchUpdate().withState(AuthenticationState.ENABLED)),
            com.azure.core.util.Context.NONE);
    }
}
