
import com.azure.resourcemanager.horizondb.models.AuthenticationState;
import com.azure.resourcemanager.horizondb.models.PasswordAuthenticationForPatchUpdate;
import com.azure.resourcemanager.horizondb.models.PasswordAuthenticationPropertiesForPatchUpdate;

/**
 * Samples for PasswordAuthentications Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/PasswordAuthentications_Enable.json
     */
    /**
     * Sample code: Enable password authentication and reset the existing administrator password.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void enablePasswordAuthenticationAndResetTheExistingAdministratorPassword(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.passwordAuthentications().update("example-resource-group", "example-cluster",
            new PasswordAuthenticationForPatchUpdate().withProperties(
                new PasswordAuthenticationPropertiesForPatchUpdate().withState(AuthenticationState.ENABLED)
                    .withAdministratorLoginPassword("fakeTokenPlaceholder")),
            com.azure.core.util.Context.NONE);
    }
}
