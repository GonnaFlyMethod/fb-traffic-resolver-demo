import { observer } from "mobx-react-lite";
import * as yup from "yup";
import { useTranslation } from "react-i18next";
import { useFormik } from "formik";
import { Button, Grid, TextField } from "@mui/material";

import {
  emailValidation,
  commonStringValidation,
} from "shared/validations";
import { TUser } from "shared/types";

import { UsersModel } from "../../model";
import styles from "./Styles.module.scss";

const validationSchema = yup.object().shape({
  name: commonStringValidation("Name", 3),
  surname: commonStringValidation("Surname", 3),
  email: emailValidation(),
});

function UpdateUserForm({ user, hideModal }: any) {
  const { id, name, surname, email } = user;

  const { t } = useTranslation();

  const { handleSubmit, values, handleChange, touched, errors } =
    useFormik<TUser>({
      initialValues: {
        id,
        name,
        surname,
        email,
      },
      validationSchema,
      onSubmit: (value: TUser) => {
        console.log(value);
        UsersModel.update(value);
        hideModal();
      },
    });

  return (
    <form onSubmit={handleSubmit} className={styles.centered}>
      <Grid item container spacing={2} direction="column">
        <Grid item>
          <TextField
            fullWidth
            id="name"
            name="name"
            label={t("user:name")}
            value={values.name}
            onChange={handleChange}
            error={touched.name && Boolean(errors.name)}
            helperText={touched.name && errors.name}
          />
        </Grid>
        <Grid item>
          <TextField
            fullWidth
            id="surname"
            name="surname"
            label={t("user:surname")}
            value={values.surname}
            onChange={handleChange}
            error={touched.surname && Boolean(errors.surname)}
            helperText={touched.surname && errors.surname}
          />
        </Grid>
        <Grid item>
          <TextField
            fullWidth
            id="email"
            name="email"
            label={t("user:email")}
            value={values.email}
            onChange={handleChange}
            error={touched.email && Boolean(errors.email)}
            helperText={touched.email && errors.email}
          />
        </Grid>
        <Grid item>
          <Button color="primary" variant="contained" fullWidth type="submit">
            {t("common.confirm")}
          </Button>
        </Grid>
      </Grid>
    </form>
  );
}

export default observer(UpdateUserForm);
