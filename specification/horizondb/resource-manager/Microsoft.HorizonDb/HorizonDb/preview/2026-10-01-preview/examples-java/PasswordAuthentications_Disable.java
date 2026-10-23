
import com.azure.resourcemanager.horizondb.models.AuthenticationState;
import com.azure.resourcemanager.horizondb.models.PasswordAuthenticationForPatchUpdate;
import com.azure.resourcemanager.horizondb.models.PasswordAuthenticationPropertiesForPatchUpdate;

/**
 * Samples for PasswordAuthentications Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/PasswordAuthentications_Disable.json
     */
    /**
     * Sample code: Disable password authentication.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void disablePasswordAuthentication(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.passwordAuthentications().update("example-resource-group", "example-cluster",
            new PasswordAuthenticationForPatchUpdate().withProperties(
                new PasswordAuthenticationPropertiesForPatchUpdate().withState(AuthenticationState.DISABLED)),
            com.azure.core.util.Context.NONE);
    }
}
