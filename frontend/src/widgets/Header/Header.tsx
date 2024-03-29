import {observer} from "mobx-react-lite";
import {AppBar, IconButton, Toolbar, Typography} from "@mui/material";
import ListIcon from "@mui/icons-material/List";

interface Props {
    className?: string;
    toggleMenu: () => void;
}

const Header = ({className, toggleMenu}: Props) => {
    return (
        <AppBar position="static" className={className}>
            <Toolbar disableGutters sx={{flexGrow: 1, padding: "0 15px"}}>
                <IconButton
                    edge="start"
                    color="inherit"
                    aria-label="menu"
                    onClick={toggleMenu}
                >
                    <ListIcon/>
                </IconButton>

                <Typography sx={{flexGrow: 1}} variant="h6">
                    Demo fb-traffic-resolver
                </Typography>
            </Toolbar>
        </AppBar>
    );
};
export default observer(Header);
