
import com.azure.resourcemanager.horizondb.models.PasswordAuthenticationForPatchUpdate;
import com.azure.resourcemanager.horizondb.models.PasswordAuthenticationPropertiesForPatchUpdate;

/**
 * Samples for PasswordAuthentications Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/PasswordAuthentications_ResetPassword.json
     */
    /**
     * Sample code: Reset the sole PostgreSQL administrator role password.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        resetTheSolePostgreSQLAdministratorRolePassword(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.passwordAuthentications().update("example-resource-group", "example-cluster",
            new PasswordAuthenticationForPatchUpdate()
                .withProperties(new PasswordAuthenticationPropertiesForPatchUpdate()
                    .withAdministratorLoginPassword("fakeTokenPlaceholder")),
            com.azure.core.util.Context.NONE);
    }
}
