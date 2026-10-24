
import com.azure.resourcemanager.horizondb.models.AuthenticationState;
import com.azure.resourcemanager.horizondb.models.MicrosoftEntraAuthenticationForPatchUpdate;
import com.azure.resourcemanager.horizondb.models.MicrosoftEntraAuthenticationPropertiesForPatchUpdate;

/**
 * Samples for MicrosoftEntraAuthentications Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/MicrosoftEntraAuthentications_Disable.json
     */
    /**
     * Sample code: Disable Microsoft Entra authentication without modifying existing database roles.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void disableMicrosoftEntraAuthenticationWithoutModifyingExistingDatabaseRoles(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.microsoftEntraAuthentications().update("example-resource-group", "example-cluster",
            new MicrosoftEntraAuthenticationForPatchUpdate().withProperties(
                new MicrosoftEntraAuthenticationPropertiesForPatchUpdate().withState(AuthenticationState.DISABLED)),
            com.azure.core.util.Context.NONE);
    }
}
